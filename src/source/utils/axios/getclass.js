import axios from "axios";
import { ActionCreator } from "../redux/action";
import { store } from "../redux/store";

export function getClass(token){

    axios({
        method:'post',
        url:'http://localhost:1234/getclass',
        headers:{
            authorization : token
        }
    })
    .then((data)=>{
        console.log(data)
        const action=ActionCreator('class',data.data.message);
        store.dispatch(action);
    })
    .catch((err)=>{
        console.log(err)
        const action = ActionCreator('logout');
        store.dispatch(action);
    })
}