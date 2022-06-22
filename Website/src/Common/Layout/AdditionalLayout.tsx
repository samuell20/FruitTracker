import {Footer} from '../Footer';
import { Header } from '../Header/Default';

export default function Layout({ children }: any) {
    return (
      <div className="d-flex h-100 flex-column">
        <Header />
        <main className="h-100 d-flex align-items-center justify-content-center">{children}</main>
        <Footer />
      </div>
    )
}