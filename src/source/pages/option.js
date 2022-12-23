import { useSelector } from "react-redux"
import { Navigate, useNavigate } from "react-router-dom";

export const WorkOption =()=>{

    const Nav=useNavigate();

    const loggedin=useSelector(state=>state.loggedin);

    if(!loggedin){
        <Navigate to='/'></Navigate>
    }
    else{
        return(
            <div className="btn-group row">
                <button className="btn btn-warning col-md-5 col-sm-11 m-1" style={{'height':'50vh'}} onClick={()=>Nav('/qr')}><h1 className="display-1">Generate QR</h1></button>
                <button className="btn btn-warning col-md-5 col-sm-11 m-1" style={{'height':'50vh'}} onClick={()=>Nav('/attend')}><h1 className="display-1">See Attendance</h1></button>
            </div>
        )
    }
    
}

