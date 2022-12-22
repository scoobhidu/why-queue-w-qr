const express= require('express');
const app= express();  
var cors = require('cors')
const useRouter=require('./source/routes/unproutes.js')
const pRouter=require('./source/routes/proutes')

app.use(cors({
    origin:'http://localhost:3000'
}))
app.use(express.urlencoded());
app.use(express.json()); 

app.use('/',useRouter);
app.use(require('./source/utils/middlewares/verifytoken'))
app.use('/',pRouter);



const server = app.listen(1234,(err)=>{
    if(err){
        console.log("Server Crash");
    }
    else{
        console.log("Server Started...",server.address().port)
    }
}) //starts our app
