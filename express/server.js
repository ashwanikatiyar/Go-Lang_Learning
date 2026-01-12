import express from 'express'

const users = [
    {id : 1 , name : "User1"},
    {id : 2 , name : "User1"},
    {id : 3 , name : "User1"},
    {id : 4 , name : "User1"},
    {id : 5 , name : "User1"},
    {id : 6 , name : "User1"},
    {id : 7 , name : "User1"},
    {id : 8 , name : "User1"},
    {id : 9 , name : "User1"},
    {id : 10 , name : "User1"},
    {id : 11 , name : "User1"},
    {id : 12 , name : "User1"},
    {id : 13 , name : "User1"},
    {id : 14 , name : "User1"},

]
const app = express()
app.use(express.json())


app.get(('/users') , (req , res)=> {
    const {limit , page} = req.query

    const statIndex = (page -1 )* limit
    const endIndex = page * limit

    const resultUsers = users.slice(statIndex , endIndex)
    res.status(200).json(resultUsers)
})

app.listen(3000)