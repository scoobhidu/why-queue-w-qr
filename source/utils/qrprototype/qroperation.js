const QRCode = require('qrcode');
const { add } = require('../../database/mongodb/operations/operation');
const { generateClassToken } = require('../jwt/jwtoperations');

const QR = {
    generator(userid){
        const qrtoken = generateClassToken();
        const presentdate=Date.now()
        return(
            QRCode.toDataURL('http://randomgoapi/'+qrtoken+'/'+presentdate)
            .then(url => {
                // console.log('HERE',url);
                add(qrtoken,userid,presentdate);
                return(url)
            })
            .catch(err => {
                console.error(err)
            })
        )
        
    }
}

module.exports=QR;
