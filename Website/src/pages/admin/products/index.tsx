import type { NextPage } from 'next'
import Link from 'next/link';
import { useState, useEffect, JSXElementConstructor, Key, ReactElement, ReactFragment, ReactPortal } from 'react';
import { productService } from 'Services/product.service';

const Products: NextPage = () => {
   const [products, setproducts] = useState<any>([]);

    useEffect(() => {
        productService.getAll().then(x => setproducts(x));
    }, []);

    function deleteproduct(id:Number) {

        setproducts(products.map((x: { id: Number; isDeleting: boolean; }) => {
            if (x.id === id) { x.isDeleting = true; }
            return x;
        }));
        productService.delete(id).then(() => {
            setproducts((products: any[]) => products.filter((x: { id: Number; }) => x.id !== id));
        });
    }

    return (
        <div>
            <h1>Productos</h1>
            
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th style={{ width: '20%' }}>Nombre</th>
                        <th style={{ width: '40%' }}>Descripción</th>
                        <th style={{ width: '20%' }}>Precio</th>
                        <th style={{ width: '10%' }}>Tipo unidad</th>
                        <th style={{ width: '10%' }}> 
                          <Link href="products/add">
                            <a className="btn btn-sm btn-success mb-2">Añadir producto</a> 
                          </Link>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {products && products.map((product: { id: Number; name: string; description: string; price: Number ; unit: string ; isDeleting: boolean | undefined; }) =>
                        <tr key={product.id.toString()}>
                            <td>{product.name}</td>
                            <td>{product.description}</td>
                            <td>{product.price.toString()}</td>
                            <td>{product.unit}</td>
                            <td style={{ whiteSpace: 'nowrap' }}>
                                <Link href={`/admin/products/edit/${product.id}`} >
                                  <a className="btn btn-sm btn-primary mr-1">
                                    Edit
                                  </a>
                                </Link>
                                <button onClick={() => deleteproduct(product.id)} className="btn btn-sm btn-danger btn-delete-product" disabled={product.isDeleting}>
                                    {product.isDeleting 
                                        ? <span className="spinner-border spinner-border-sm"></span>
                                        : <span>Delete</span>
                                    }
                                </button>
                            </td>
                        </tr>
                    )}
                    {!products &&
                        <tr>
                            <td colSpan={4} className="text-center">
                                <div className="spinner-border spinner-border-lg align-center"></div>
                            </td>
                        </tr>
                    }
                    {products && !products.length &&
                        <tr>
                            <td colSpan={4} className="text-center">
                                <div className="p-2">No products to display</div>
                            </td>
                        </tr>
                    }
                </tbody>
            </table>
        </div>
    );
}

export default Products 

