import Link from "next/link";
import styles from '../../../../styles/Header.module.css'
import Image from 'next/image'
import logo from '../../../../public/logo.png'
const Header= () =>{
   return (
       <header>
            <div className={styles.sides}>
               <Image src={logo} width={70} height={40}/> 
            </div>
            <nav className={styles.navbar}>
                <ul className={styles.links}>
                    <li><Link href="/">Inicio</Link></li>
                    <li><Link href="/store/products"> Productos </Link></li> 
                    <li><Link href="/store/about">Sobre nosotros</Link></li>
                    <li><Link href="/store/contact">Contacto</Link></li>
                </ul>
            </nav>
            <div className={styles.sides}>
               <a href="/store/login"><i className="bi bi-person fs-3 text-light"></i></a> 
            </div>
       </header>
   ) 
}


export default Header;