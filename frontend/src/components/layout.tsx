import React from 'react';
import { Link } from 'react-router-dom';

type LayoutProps = {
  children: React.ReactNode
}

function Layout({ children }: LayoutProps) {
  return (
    <>
      <header>
        <nav className="navbar navbar-expand-md navbar-light">
          <div className="container">
            <Link to="/" className="navbar-brand" href="#">Codenv</Link>
            <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
              <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarCollapse">
              <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                  <Link to="/" className="nav-link active" aria-current="page">Home</Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>
      </header>

      <main className="flex-shrink-0 mb-5">
        <div className="container pt-5">
          { children }
        </div>
      </main>

      <footer className="footer mt-auto py-5 bg-light">
        <div className="container">
          <span className="text-muted">Place sticky footer content here.</span>
        </div>
      </footer>

    </>
  )
}

export default Layout;