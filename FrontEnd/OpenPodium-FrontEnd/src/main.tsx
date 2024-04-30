// Base Imports
import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'

// Extra Imports
import {
  createBrowserRouter,
  RouterProvider
} from "react-router-dom"

// Page Imports
import App from './App.tsx'
import Login from "./Pages/Login/Login.tsx"
import Home from './Pages/Home/Home.tsx'

// Initialise Pages
const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: []
  },
  {
    path: "/login",
    element: <Login />,
    children: []
  },
  {
    path: "/home",
    element: <Home />,
    children: []
  }
])

// Render the pages
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <div id="main-wrapper">
      <RouterProvider router={router} />
    </div>
  </React.StrictMode>
)
