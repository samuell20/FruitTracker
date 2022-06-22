import type { NextPage } from 'next'
import Link from 'next/link';
import { useState, useEffect, JSXElementConstructor, Key, ReactElement, ReactFragment, ReactPortal } from 'react';
import { userService } from 'Services';

const Users: NextPage = () => {
   const [users, setUsers] = useState<any>([]);

    useEffect(() => {
        userService.getAll().then(x => setUsers(x));
    }, []);

    function deleteUser(id:Number) {

        setUsers(users.map((x: { id: Number; isDeleting: boolean; }) => {
            if (x.id === id) { x.isDeleting = true; }
            return x;
        }));
        userService.delete(id).then(() => {
            setUsers((users: any[]) => users.filter((x: { id: Number; }) => x.id !== id));
        });
    }

    return (
        <div>
            <h1>Users</h1>
            
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th style={{ width: '30%' }}>Nombre usuario</th>
                        <th style={{ width: '30%' }}>Email</th>
                        <th style={{ width: '30%' }}>Rol</th>
                        <th style={{ width: '10%' }}> 
                          <Link href="/admin/users/add">
                            <a className="btn btn-sm btn-success mb-2">Add User</a> 
                          </Link>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {users && users.map((user: { id: Number; username: string; email: string ; role: string ; isDeleting: boolean | undefined; }) =>
                        <tr key={user.id.toString()}>
                            <td>{user.username} </td>
                            <td>{user.email}</td>
                            <td>{user.role}</td>
                            <td style={{ whiteSpace: 'nowrap' }}>
                                <Link href={`/admin/users/edit/${user.id}`} >
                                  <a className="btn btn-sm btn-primary mr-1">
                                    Edit
                                  </a>
                                </Link>
                                <button onClick={() => deleteUser(user.id)} className="btn btn-sm btn-danger btn-delete-user" disabled={user.isDeleting}>
                                    {user.isDeleting 
                                        ? <span className="spinner-border spinner-border-sm"></span>
                                        : <span>Delete</span>
                                    }
                                </button>
                            </td>
                        </tr>
                    )}
                    {!users &&
                        <tr>
                            <td colSpan={4} className="text-center">
                                <div className="spinner-border spinner-border-lg align-center"></div>
                            </td>
                        </tr>
                    }
                    {users && !users.length &&
                        <tr>
                            <td colSpan={4} className="text-center">
                                <div className="p-2">No Users To Display</div>
                            </td>
                        </tr>
                    }
                </tbody>
            </table>
        </div>
    );
}

export default Users 

