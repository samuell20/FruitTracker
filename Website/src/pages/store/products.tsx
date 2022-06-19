import type { NextPage } from 'next'
import img from '../../../public/Fruits.jpg'
import Image from 'next/image'

const Products: NextPage = ({product_data}:any) => {
    const products = product_data.map((item : any) => (
        <div className="card w-25 mb-3 mx-2" key={item.id}>
            
            <Image src={img} className="card-img-top"/> 
            <div className="card-body">
                <h5 className="card-title">{item.name}</h5>
                <p className="card-text">
                    {item.description} <br /> 
                    {item.price} {item.currency} <br />    
                </p>
              </div>

        </div>

    ))
  return (
     <div className="main-content">
       <h1>Productos</h1>
       <div className="d-flex flex-wrap justify-content-between">
         {products}
       </div>
     </div> 
  )
}

export async function getStaticProps() {
  const res = await fetch('http://127.0.0.1:4000/products')
  const product_data = await res.json()
  return {
    props: {
      product_data,
    },
    // Next.js will attempt to re-generate the page:
    // - When a request comes in
    // - At most once every 10 seconds
    revalidate: 10, // In seconds
  }
}
export default Products