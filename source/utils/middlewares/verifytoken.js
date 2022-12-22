const { decodeLoginToken } = require("../jwt/jwtoperations");

const tokenVerify = (request, response, next)=>{
    
    const tokenId = request.headers['authorization']; 
    console.log('Token Rec in Middleware ', tokenId);
     const isVerified = decodeLoginToken(tokenId); 
     console.log('Verify Status',isVerified);
     if(isVerified){
        console.log('Verified');
        next();
     }
     else{
        console.log('Not Verified');
        response.status(401).json({message:'UnAuthorized User'});
     }
}
module.exports = tokenVerify;