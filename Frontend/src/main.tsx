import React from 'react'
import ReactDOM from 'react-dom/client'
import InputArea from './InputArea'
import NavBarPrincipal from './NavBarPrincipal';
import './index.css'



ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>

        <NavBarPrincipal/>
        <InputArea/>
    </React.StrictMode>,
)
