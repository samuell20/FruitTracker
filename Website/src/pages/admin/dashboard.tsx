import type { NextPage } from 'next'
import Link from 'next/link'

const Dashboard: NextPage = () => {
  return (
     <div className="h-100">
       <h1>Panel de control</h1> 
        <div className="w-50 h-100 m-auto d-flex flex-wrap justify-content-around">
            <Link href="users">
              <a className="border w-50 h-25 d-flex justify-content-center align-items-center flex-column">
                <h5>Usuarios</h5>
                <p>Pantalla para ver o editar los usuarios</p>
              </a>
            </Link>
            <Link href="products">
              <a className="border w-50 h-25 d-flex justify-content-center align-items-center flex-column">
                <h5>Productos</h5>
                <p>Pantalla para ver o editar los productos</p>
              </a>
            </Link>
            <Link href="orders">
              <a className="border w-50 h-25 d-flex justify-content-center align-items-center flex-column">
                <h5>Pedidos</h5>
                <p>Pantalla para ver o editar los productos</p>
              </a>
            </Link>
            <Link href="tickets">
              <a className="border w-50 h-25 d-flex justify-content-center align-items-center flex-column">
                <h5>Tickets</h5>
                <p>Pantalla para ver los tickets</p>
              </a>
            </Link>
            <Link href="invoices">
              <a className="border w-50 h-25 d-flex justify-content-center align-items-center flex-column">
                <h5>Facturas</h5>
                <p>Pantalla en la que podrá ver las diferentes facturas</p>
              </a>
            </Link>
        </div> 
     </div> 
  )
}

export default Dashboard