import { useRouter } from 'next/router';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as Yup from 'yup';

import Link from 'next/link';
import { productService,alertService } from 'Services';

export function AddEdit(props:any) {
    const product = props?.params.product;
    const isAddMode = !product;
    const router = useRouter();
    
    // form validation rules 
    const validationSchema = Yup.object().shape({ 
        name: Yup.string()
            .required('Debe proporcionar un nombre'), 
        description: Yup.string()
            .required('Debe proporcionar una descripción'),
        unit: Yup.string()
            .required('Debe proporcionar un rol'),
        price: Yup.number()
            .required('Debe proporcionar un precio')
            
    });
    const formOptions = { resolver: yupResolver(validationSchema), defaultValues:{} };

    // set default form values if product passed in props
    if (!isAddMode) {
        const { ...defaultValues } = product;
        formOptions.defaultValues = defaultValues;
    }

    // get functions to build form with useForm() hook
    const { register, handleSubmit, reset, formState } = useForm(formOptions);
    const { errors } = formState;

    function onSubmit(data:any) {
        console.log("submiting")
        return isAddMode
            ? createproduct(data)
            : updateproduct(product.id, data);
    }

    function createproduct(data:any) {
        return productService.create(data)
            .then(() => {
                alertService.success('Product added', { keepAfterRouteChange: true });
                router.push('.');
            })
            .catch(alertService.error);
    }

    function updateproduct(id:Number, data:object) {
        return productService.update(id, data)
            .then(() => {
                alertService.success('Product updated', { keepAfterRouteChange: true });
                router.push('..');
            })
            .catch(alertService.error);
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="w-75 m-auto">
            <h1>{isAddMode ? 'Añadir producto' : 'Editar usuario'}</h1>
            <div className="row mt-3"> 
                <div className="form-group col-5">
                    <label>Nombre del producto</label>
                    <input  type="text" {...register('name')} className={`form-control ${errors.name ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.name?.message}</div>
                </div>
                <div className="form-group col">
                    <label>Medida</label>
                    <select  {...register('unit')} className={`form-control ${errors.unit ? 'is-invalid' : ''}`}>
                        <option value=""></option>
                        <option value="Unidad">Unidad</option>
                        <option value="Kg">Kg</option>
                    </select>
                    <div className="invalid-feedback">{errors.unit?.message}</div>
                </div>
            </div>
            <div className="row mt-3"> 
                <div className="form-group col-10">
                    <label>Descripción del producto</label>
                    <input  type="text" {...register('description')} className={`form-control ${errors.description ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.description?.message}</div>
                </div> 
                <div className="form-group col-2">
                    <label>Precio</label>
                    <input  type="text" {...register('price')} className={`form-control ${errors.price ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.price?.message}</div>
                </div> 
            </div> 
            <div className="form-group mt-3">
                <button type="submit" disabled={formState.isSubmitting} className="btn btn-primary me-2">
                    {formState.isSubmitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                    Guardar
                </button>
                <button onClick={() => reset(formOptions.defaultValues)} type="button" disabled={formState.isSubmitting} className="btn btn-secondary me-2">Resetear</button>
                <Link href="/admin/products" className="btn btn-link">Cancelar</Link>
            </div>
        </form>
    );
}