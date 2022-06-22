import React from 'react'

function Footer() {
  return (
    <footer className="mt-auto">
      <div className="d-flex justify-content-between align-items-center h-100 w-70">
        <p className="m-0 p-0">Copyright &copy 2022 | FruitTracker </p>
        <div>
          <a href="#">
            <i className="bi bi-instagram fs-6 me-2"></i>
          </a>
          <a href="#">
            <i className="bi bi-facebook fs-6"></i>
          </a>
        </div>
      </div> 
    </footer>
  )
}

export default Footer