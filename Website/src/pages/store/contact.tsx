import type { NextPage } from 'next'
import Image from 'next/image'
import FruitImg from '../../../public/Fruits.jpg'
import Styles from '../../../styles/store/Contact.module.css'

const Contact: NextPage = () => {
  return (
     <div className="main-content">
      <h1>Contacto</h1>
      <div className="d-flex justify-content-center align-items-center">
        <Image src={FruitImg} alt="Cesta de frutas" width={600} height={400}/>
      </div>
      <div className="d-flex justify-content-center align-items-center mt-5 mb-5">
        <div className={Styles.tags + ' me-5'}>
          <h3>Lorem Ipsum</h3>
        </div>
        <div className={Styles.tags}>
          <h3>Lorem Ipsum</h3>
        </div>
      </div>
      <h2 className="text-center mb-5">Nuestras tiendas</h2>
      <div className="d-flex justify-content-between align-items-center">
        <div className="me-5">
          <Image src={FruitImg} alt="Cesta de frutas" />
          <h4 className="text-center">Lorem ipsum</h4>
        </div>
        <div>
          <Image src={FruitImg} alt="Cesta de frutas" />
          <h4 className="text-center">Lorem ipsum</h4>
        </div>
      </div>
     </div> 
  )
}

export default Contact