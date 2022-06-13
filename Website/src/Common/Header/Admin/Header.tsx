import Link from "next/link";
import styles from '../../../../styles/Header.module.css'

const Header= () =>{
   return (
       <header>
            <div className={styles.sides}>
                
            </div>
            
            <div className={styles.sides}>
               <a href=""><i className="bi bi-person fs-3 text-light"></i></a> 
            </div>
       </header>
   ) 
}


export default Header;