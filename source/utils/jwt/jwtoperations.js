const jwt = require('jsonwebtoken');
require('dotenv').config()

const obj = {
    SECRET: 'Hello',
    generateLoginToken(userid,password){
        console.log(process.env.TEACHER_TOKEN);
        const token = jwt.sign({'userName':userid,'password':password},process.env.TEACHER_TOKEN,{
            //algorithm:'RS256',
            expiresIn:'1200s'
        });
        // console.log('Token ');
        return token;
    },
    decodeLoginToken(tokenId){
        try{
            const decode = jwt.verify(tokenId,process.env.TEACHER_TOKEN);
            console.log(decode);
            if(decode && decode.userName){
                console.log(decode.userid);
                return true;
            }
            else{
                console.log("Invalid Token ");
                return false;
            }
        }
        catch(err){
            console.log('Catch Error in Decode::::: ', err);
            return false;
        }
    },
    generateClassToken(teacher){
        console.log(process.env.ATTENDANCE_TOKEN);
        const token = jwt.sign({'userName':teacher},process.env.TEACHER_TOKEN,{
            //algorithm:'RS256',
            expiresIn:'12s'
        });
        // console.log('Token Attendance');
        return token;
    }

}
module.exports = obj;
