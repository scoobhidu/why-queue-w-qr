import axios from "axios";

export function Verify(token){
    axios({
        method:'post',
        url:'http://localhost:1234/verify',
        headers:{
            'authorization': token
        }
    })
    .then((data)=>{
        console.log('Verify',data);
        return true;
    })
    .catch((err)=>{
        console.log(err);
        return false;
    })
}