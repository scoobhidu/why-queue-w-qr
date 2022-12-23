import { useSelector } from "react-redux"
import { Navigate } from "react-router-dom"
import { LoginValidation } from "../components/login"

export const Login=()=>{

    const logginstatus = useSelector(state=>state.loggedin)

    if(logginstatus){
        return(<Navigate to='/classes'></Navigate>)
    }
    else{
        return(
            <div className="bg-purple-300" style={{'minHeight':'100vh','minWidth':'100vw'}}>
                <div className="m-auto w-4/5" 
                    style={{
                    //     'minHeight': '80vh',
                    //     'width': '60%',
                    //     'background': 'linear-gradient(to right bottom,rgba(255, 255, 255, 0.7),rgba(255, 255, 255, 0.3))',
                    //     'borderRadius': '2rem',
                        'zIndex': '2',
                        'backdropFilter': 'blur(2rem)',
                        'display': 'flex',
                        'padding':'0px'
                        }}
                    >
                    <div className="d-flex align-items-center w-full">
                        <LoginValidation/>
                    </div>
                    
                </div>
                
            </div>
        )
    }

    
}