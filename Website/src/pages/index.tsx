import type { NextPage } from 'next'
import Link from 'next/link'
import Image from 'next/image'
import FruitImg from '../../public/Fruits.jpg'
import Styles from '../../styles/Home.module.css'; 
import Fruit1 from '../../public/fresa.jpg'
import Fruit2 from '../../public/coco.jpg'
import Fruit3 from '../../public/naranja.jpg'

const Home:NextPage = () => {
  
  return (
    <>
    <div className="main-content">
       <h4 className="mb-5">Inicio</h4>
        <section className="d-flex w-100">
        <div className="w-50">
          <Image src={FruitImg}/>
        </div>
        <div className="d-flex justify-content-center align-items-center w-50 p-5">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
              Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 
              fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa 
              qui officia deserunt mollit anim id est laborum.</p>
          </div>
        </section>
     </div> 
     <section className={Styles.section} >
          <div className="d-flex align-items-center justify-content-center h-100">
            <div className="border h-50 w-15 p-3 d-flex align-items-center justify-content-center me-3">
              <p className="me-2 m-0 p-0">
                Pedidos a domicilio
              </p> 
              <i className="bi bi-truck fs-1 text-green-p"></i>
            </div>
            <div className="border h-50 w-15 p-3 d-flex align-items-center justify-content-center">
              <p className="me-2 m-0 p-0">
                Pedidos en tiendas
              </p>
              <i className="bi bi-shop fs-1 text-green-p"></i>
            </div>
            <div className="border h-50 w-15 p-3 d-flex align-items-center justify-content-center ms-3">
              <p className="me-2 m-0 p-0">
                Los mejores productos
              </p>
              <i className="bi bi-basket fs-1 text-green-p"></i>
            </div>
          </div>
      </section> 
      <section className="main-content mb-5">
    <h3 className="text-center mb-5">Productos más vendidos</h3>
        <div className="d-flex justify-content-between">
          <Image src={Fruit1} width={400} height={400} />
          <Image src={Fruit2} width={400} height={400}/>
          <Image src={Fruit3} width={400} height={400}/>
        </div>
      </section>
    </>
  )
}

export default Home
