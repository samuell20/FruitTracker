import type { NextPage } from 'next'
import Link from 'next/link'
import Image from 'next/image'
import FruitImg from '../../public/Fruits.jpg'
import Styles from '../../styles/Home.module.css'; 

const Home:NextPage = () => {
  
  return (
    <>
    <div className="main-content">
       <h3>Inicio</h3>
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

      </section> 
      <section className="main-content mt-5 mb-5">
    <h3 className="text-center mb-5">Productos más vendidos</h3>
        <div className="d-flex justify-content-between">
          <div className={Styles.products}></div>
          <div className={Styles.products}></div>
          <div className={Styles.products}></div>
        </div>
      </section>
    </>
  )
}

export default Home
