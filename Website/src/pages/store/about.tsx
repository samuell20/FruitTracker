import type { NextPage } from 'next'
import Image from 'next/image'
import FruitImg from '../../../public/Fruits.jpg'
import Styles from '../../../styles/store/About.module.css'

const About: NextPage = () => {
  return (
    <>
      <section className="main-content">
        <h3>Sobre nosotros</h3>
        <div>
          <h2>¿Quienes somos?</h2>
          <div className="d-flex">
              <div className="w-50">
                <Image src={FruitImg}/>
              </div>
              <div className="d-flex justify-content-center align-items-center w-50 p-5">
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
                  incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
                  exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
                  Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 
                  fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa 
                  qui officia deserunt mollit anim id est laborum.
                </p>
            </div>
          </div>
        </div> 
      </section> 
      <section className={Styles.section + ' d-flex text-light mt-5'}>
        <div className="d-flex w-80 mt-3 mb-3">
          <div className="w-50 d-flex flex-column justify-content-center align-items-center">
            <h3 className="text-center">Nuestros productos</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
              Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 
              fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa 
              qui officia deserunt mollit anim id est laborum.
            </p>
          </div> 
          <div className="w-50 d-flex justify-content-center align-items-center">
            <Image src={FruitImg} width={400} height={300}/> 
          </div>
        </div>
      </section>
      <section className="main-content">
        <h3 className="text-center mb-5" >Nuestras variedades</h3>
        <div className="d-flex justify-content-between align-items-center">
          <div className="w-25">
            <h4 className="text-center">Frutas</h4>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
            </p>
          </div>
          <div className="w-25">
            <h4 className="text-center">Verduras</h4>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
            </p>
          </div>
          <div className="w-25">
            <h4 className="text-center">Otros productos</h4>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
            </p> 
          </div>
        </div>
      </section>
    </>
  )
}

export default About
