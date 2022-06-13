import '../../styles/globals.css'
import WebLayout from '../Common/Layout/WebLayout'
import AdminLayout from '../Common/Layout/AdminLayout'
import { useRouter } from 'next/router'


export default function MyApp({ Component, pageProps }: any ) {
  // Use the layout defined at the page level, if available
  const router = useRouter()
  let Layout = WebLayout  
  if(router.pathname.includes('/admin'))
  {
    Layout = AdminLayout
  }
  
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout> 
  ) 
}
