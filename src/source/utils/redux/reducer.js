
export const reducerfn=(state={
    loggedin:false,
    token:'',
    classes:[],
    QR:''
},action)=>{
    if(action.type==='login'){
        let token=action.payload;
        let loggedin=true;
        return{...state,loggedin:loggedin,token:token}
    }
    else if(action.type==='logout'){
        const token='';
        return{...state,loggedin:false,token:token}
    }
    else if(action.type==='class'){
        if(state.classes===action.payload){
            return{...state};
        }
        const classes=action.payload;
        return{...state,classes:classes}
    }
    else if(action.type==='QR'){
        const src=action.payload;
        return {...state,QR:src}
    }
    else{
        return state;
    }
}