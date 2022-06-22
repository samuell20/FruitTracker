import Link from "next/link";

const Sidebar = () => {

    return (
    <div className="col-auto col-md-3 col-xl-2 px-sm-2 px-0 bg-dark">
        <div className="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
            <Link href="dashboard" className="d-flex align-items-center pb-3 mb-md-0 me-md-auto text-white text-decoration-none">
                <span className="fs-5 d-none d-sm-inline">Menú</span>
            </Link>
            <ul className="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
                <li className="nav-item">
                    <Link href="/admin/dashboard" >
                        <a className="nav-link px-0 align-middle">
                            <i className="fs-4 bi-speedometer2"></i> <span className="ms-1 d-none d-sm-inline">Panel de control</span> 
                        </a>
                    </Link> 
                </li>
                <li className="nav-item">
                    <Link href="/admin/users">
                        <a className="nav-link align-middle px-0">
                            <i className="fs-4 bi bi-people"></i> <span className="ms-1 d-none d-sm-inline">Usuarios</span>
                        </a>
                    </Link>
                </li> 
                <li className="nav-item">
                    <Link href="/admin/products">
                        <a className="nav-link px-0 align-middle ">
                            <i className="fs-4 bi bi-box"></i> <span className="ms-1 d-none d-sm-inline">Productos</span> 
                        </a>
                    </Link> 
                </li>
                <li className="nav-item">
                    <Link href="/admin/orders" >
                        <a className="nav-link align-middle px-0">
                            <i className="fs-4 bi bi-card-checklist"></i> <span className="ms-1 d-none d-sm-inline">Pedidos</span>
                        </a>
                    </Link>
                </li>  
                <li className="nav-item">
                    <Link href="/admin/tickets">
                        <a className="nav-link px-0 align-middle">
                            <i className="fs-4 bi bi-ticket-detailed"></i> <span className="ms-1 d-none d-sm-inline">Tickets</span>
                        </a>
                    </Link>
                </li>
                <li className="nav-item">
                    <Link href="/admin/invoices" >
                        <a className="nav-link px-0 align-middle ">
                            <i className="fs-4 bi bi-receipt"></i> <span className="ms-1 d-none d-sm-inline">Facturas</span>
                        </a>
                    </Link> 
                </li>  
            </ul>
            <hr />
            <div className="dropdown pb-4">
                <a href="#" className="d-flex align-items-center text-white text-decoration-none dropdown-toggle" id="dropdownUser1" data-bs-toggle="dropdown" aria-expanded="false">
                    <span className="d-none d-sm-inline mx-1">loser</span>
                </a>
                <ul className="dropdown-menu dropdown-menu-dark text-small shadow">
                    <li><a className="dropdown-item" href="#">New project...</a></li>
                    <li><a className="dropdown-item" href="#">Settings</a></li>
                    <li><a className="dropdown-item" href="#">Profile</a></li>
                    <li>
                        <hr className="dropdown-divider" />
                    </li>
                    <li><a className="dropdown-item" href="#">Sign out</a></li>
                </ul>
            </div>
        </div>
    </div>   
    )    
}

export default Sidebar