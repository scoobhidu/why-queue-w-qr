const express=require('express');
const router=express.Router();
const QR= require('../utils/qrprototype/qroperation')
require('dotenv').config()

// const QRCode = require('qrcode')

router.get("/genqr",async (req,res)=>{
    // console.log('QR',QR.QR.generator());
    // res.send()
    const QRgen=await QR.generator();
    console.log('QR',QRgen);
    res.status(202).json({message:QRgen});
})

router.post('/getclass',(req,res)=>{ 
    //need to fetch class of teacher according to criteria mentioned in discord notes 
    // also mention what all details need to be passed to get required data
    // send array of class as written below
    res.status(202).json({message:['C1','C2','C3','C4','C1','C2','C3','C4','C1','C2','C3','C4']})
})

module.exports=router