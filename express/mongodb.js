const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');

const app = express();
app.use(cors()); // Enable CORS
app.use(express.json());

const mongoURI = 'mongodb://localhost:27017/testdb';

mongoose.connect(mongoURI)
  .then(() => console.log('MongoDB connected'))
  .catch(err => console.log(err));

const ItemSchema = new mongoose.Schema({
  name: String,
  description: String,
});
const Item = mongoose.model('Item', ItemSchema);

app.get('/items', async (req, res) => {
  try {
    const items = await Item.find();
    res.json(items);
  } catch (err) {
    res.status(500).json({ error: 'Server Error' });
  }
});

app.post('/items', async (req, res) => {
  const { name, description } = req.body;
  try {
    const newItem = new Item({ name, description });
    await newItem.save();
    res.json(newItem);
  } catch (err) {
    res.status(500).json({ error: 'Server Error' });
  }
});

const PORT = 1000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));