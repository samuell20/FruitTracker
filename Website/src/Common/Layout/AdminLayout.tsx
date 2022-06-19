import {Sidebar} from '../Sidebar'

export default function Layout({ children }: any) {
    return (
    <div className="container-fluid">
        <div className="row flex-nowrap">
            <Sidebar />   
            <div className="col p-2 m-3 bg-light">
                {children} 
            </div>
        </div>
    </div>
      
    )
}


