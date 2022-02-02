
function Header() {
    return (
        <div>
            <nav className="navbar navbar-expand-lg navbar-dark ftco_navbar bg-dark ftco-navbar-light" id="ftco-navbar">
                <div className="container">
                    <h1 className="navbar-brand">Easy Lease</h1>
                    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#ftco-nav"
                        aria-controls="ftco-nav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="oi oi-menu"></span> Menu
                    </button>
                    <div className="collapse navbar-collapse" id="ftco-nav">
                        <ul className="navbar-nav ml-auto">
                            <li className="nav-item active">
                                <a href="https://preview.colorlib.com/theme/findstate/index.html"
                                    className="nav-link">Home</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/about.html"
                                    className="nav-link">About</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/services.html"
                                    className="nav-link">Services</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/agent.html"
                                    className="nav-link">Agent</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/properties.html"
                                    className="nav-link">Listing</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/blog.html"
                                    className="nav-link">Blog</a></li>
                            <li className="nav-item">
                                <a href="https://preview.colorlib.com/theme/findstate/contact.html"
                                    className="nav-link">Contact</a></li>
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
    )
}

export default Header;