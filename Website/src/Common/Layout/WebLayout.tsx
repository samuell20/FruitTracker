import {Footer} from '../Footer';
import { Header } from '../Header/Default';


export default function Layout({ children }: any) {
    return (
      <>
        <Header />
        <main>{children}</main>
        <Footer />
      </>
    )
}