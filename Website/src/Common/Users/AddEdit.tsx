import { useRouter } from 'next/router';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as Yup from 'yup';

import Link from 'next/link';
import { userService, alertService } from 'Services';


export function AddEdit(props:any) {
    const user = props?.params.user;
    const isAddMode = !user;
    const router = useRouter();
    const [showPassword, setShowPassword] = useState(false);
    
    // form validation rules 
    const validationSchema = Yup.object().shape({ 
        username: Yup.string()
            .required('Debe proporcionar un nombre'), 
        email: Yup.string()
            .email('El email no es válido')
            .required('Debe proporcionar un email'),
        role: Yup.string()
            .required('Debe proporcionar un rol'),
        password: Yup.string()
            .transform(x => x === '' ? undefined : x)
            .concat(isAddMode ? Yup.string().required('Debe proporcionar una contraseña') : Yup.string())
            .min(6, 'La contraseña tiene que tener al menos 6 caracteres'),
        confirmPassword: Yup.string()
            .transform(x => x === '' ? undefined : x)
            .when('password', (password, schema) => {
                if (password || isAddMode) return schema.required('Debe confirmar la contraseña');
            })
            .oneOf([Yup.ref('password')], 'Las contraseñas deben coincidir')
    });
    const formOptions = { resolver: yupResolver(validationSchema), defaultValues:{} };

    // set default form values if user passed in props
    if (!isAddMode) {
        const { password, confirmPassword, ...defaultValues } = user;
        formOptions.defaultValues = defaultValues;
    }

    // get functions to build form with useForm() hook
    const { register, handleSubmit, reset, formState } = useForm(formOptions);
    const { errors } = formState;

    function onSubmit(data:any) {
        return isAddMode
            ? createUser(data)
            : updateUser(user.id, data);
    }

    function createUser(data:any) {
        return userService.create(data)
            .then(() => {
                alertService.success('User added', { keepAfterRouteChange: true });
                router.push('.');
            })
            .catch(alertService.error);
    }

    function updateUser(id:Number, data:object) {
        return userService.update(id, data)
            .then(() => {
                alertService.success('User updated', { keepAfterRouteChange: true });
                router.push('..');
            })
            .catch(alertService.error);
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="w-75 m-auto">
            <h1>{isAddMode ? 'Añadir usuario' : 'Editar usuario'}</h1>
            <div className="row mt-3"> 
                <div className="form-group col-5">
                    <label>Nombre de usuario</label>
                    <input  type="text" {...register('username')} className={`form-control ${errors.firstName ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.firstName?.message}</div>
                </div>
                <div className="form-group col-7">
                    <label>Email</label>
                    <input  type="text" {...register('email')} className={`form-control ${errors.email ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.email?.message}</div>
                </div>
            </div>
            <div className="row mt-3"> 
                <div className="form-group col">
                    <label>Rol</label>
                    <select  {...register('role')} className={`form-control ${errors.role ? 'is-invalid' : ''}`}>
                        <option value=""></option>
                        <option value="User">Usuario</option>
                        <option value="Admin">Administrador</option>
                    </select>
                    <div className="invalid-feedback">{errors.role?.message}</div>
                </div>
            </div>
            {!isAddMode &&
                <div>
                    <h3 className="pt-3">Cambiar contraseña</h3>
                    <p>Dejar en blanco para mantener la contraseña</p>
                </div>
            }
            <div className="row mt-3">
                <div className="form-group col">
                    <label>
                        Contraseña
                        {!isAddMode &&
                            (!showPassword
                                ? <span> - <a onClick={() => setShowPassword(!showPassword)} className="text-primary">Enseñar contraseña</a></span>
                                : <em> - {user.password}</em>
                            )
                        }
                    </label>
                    <input  {...register('password')} className={`form-control ${errors.password ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.password?.message}</div>
                </div>
                <div className="form-group col">
                    <label>Confirmar contraseña</label>
                    <input  type="password" {...register('confirmPassword')} className={`form-control ${errors.confirmPassword ? 'is-invalid' : ''}`} />
                    <div className="invalid-feedback">{errors.confirmPassword?.message}</div>
                </div>
            </div>
            <div className="form-group mt-3">
                <button type="submit" disabled={formState.isSubmitting} className="btn btn-primary me-2">
                    {formState.isSubmitting && <span className="spinner-border spinner-border-sm mr-1"></span>}
                    Guardar
                </button>
                <button onClick={() => reset(formOptions.defaultValues)} type="button" disabled={formState.isSubmitting} className="btn btn-secondary me-2">Reset</button>
                <Link href="/admin/users" className="btn btn-link">Cancelar</Link>
            </div>
        </form>
    );
}