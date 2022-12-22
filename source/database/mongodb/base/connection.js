const mongoose = require('mongoose');
const { URL } = require('./config');
mongoose.connect(URL,{maxPoolSize:5},(err)=>{
    if(err){
        console.log('DB Connection Problem ', err);
    }
    else{
        console.log('Connection Created...');
    }
});
module.exports= mongoose;