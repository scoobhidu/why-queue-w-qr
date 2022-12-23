import { Route, Routes } from "react-router-dom"
import { QR } from "../components/qrgenerator"
import { ClassList } from "./classlist"
import { Login } from "./login"
import { WorkOption } from "./option"

export const MainPage=()=>{
    return(
        <div className="bg-purple-300">
            <Routes>
                <Route path='/' element={<Login/>} />
                <Route path="/classes" element={<ClassList/>} />
                {/* <Route path="/choice/:class" element={<WorkOption/>} /> */}
                
                <Route path='/qr' element={<QR/>} />
            
            </Routes>
        </div>
    )
}