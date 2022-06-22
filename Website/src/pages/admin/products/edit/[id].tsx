import type { NextPage } from 'next'
import Alert from '../../../../Common/Alert'
import { AddEdit } from 'Common/Products'
import { productService } from 'Services/product.service'

const Edit: NextPage = (params) => {
  return (
     <>
      <Alert /> 
      <AddEdit params={params} />
     </> 
  )
}

export async function getServerSideProps({ params }:any) {
    const product = await productService.getById(params.id);

    return {
        props: { product}
    }
}
export default Edit 