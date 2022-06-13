import {Sidebar} from '../Sidebar'

export default function Layout({ children }: any) {
    return (
    <div className="container-fluid">
        <div className="row flex-nowrap">
         <Sidebar />   
            <div className="col py-3">
               {children} 
            </div>
        </div>
    </div>
      
    )
}


