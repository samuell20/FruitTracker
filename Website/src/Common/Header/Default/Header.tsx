import Link from "next/link";
import styles from '../../../../styles/Header.module.css'

const Header= () =>{
   return (
       <header>
            <div className={styles.sides}>
                
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
               <a href=""><i className="bi bi-person fs-3 text-light"></i></a> 
            </div>
       </header>
   ) 
}


export default Header;