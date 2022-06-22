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
                    {item.price}€ {item.unit} <br />    
                </p>
              </div>

        </div>

    ))
  return (
     <div className="main-content">
       <h6 className="mb-3">Inicio/Productos</h6>
       <h2>Catalogo de productos</h2>
       <div className="mt-5 d-flex flex-wrap justify-content-between">
         {products}
       </div>
     </div> 
  )
}

export async function getStaticProps() {
  const res = await fetch('http://fruittrackerapp.tk:80/api/products')
  const product_data = await res.json()
  console.log()
  return {
    props: {
      product_data,
    },
    
    revalidate: 10, 
  }
}
export default Products