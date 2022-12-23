import { createStore } from "redux";
import { reducerfn } from "./reducer";


export const store=createStore(reducerfn)
store.subscribe(state=>{
    console.log('State Updated... ',state);

})