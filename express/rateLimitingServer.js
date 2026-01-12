import { RateLimiterRedis, RateLimiterMemory } from 'rate-limiter-flexible';
import Redis from 'ioredis'; // Popular 2026 choice for Redis
import express from 'express';

const app = express();
const redisClient = new Redis({ /* your config */ });

// Shared Redis limiter
const redisLimiter = new RateLimiterRedis({
  storeClient: redisClient,
  keyPrefix: 'middleware',
  points: 10, 
  duration: 1,
  // Insurance: use local memory if Redis is down
  insuranceLimiter: new RateLimiterMemory({
    points: 10,
    duration: 1,
  }),
});

const rateLimiterMiddleware = (req, res, next) => {
  redisLimiter.consume(req.ip)
    .then((rateLimiterRes) => {
      // Production best practice: inform the client of their status
      res.setHeader('X-RateLimit-Limit', 10);
      res.setHeader('X-RateLimit-Remaining', rateLimiterRes.remainingPoints);
      next();
    })
    .catch((rejRes) => {
      const seconds = Math.round(rejRes.msBeforeNext / 1000) || 1;
      res.set('Retry-After', String(seconds));
      res.status(429).json({ error: 'Too Many Requests' });
    });
};

app.use(rateLimiterMiddleware);


app.get('/api/data', (req, res) => {
  res.json({ message: 'Success! Your request was within the limits.' });
});

// 6. SERVER START
const PORT = process.env.PORT || 4000;
app.listen(PORT, () => {
  console.log(`[Production Server] Listening on port ${PORT}`);
});