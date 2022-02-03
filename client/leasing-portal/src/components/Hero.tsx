function Hero() {
    return (
        <>
            <div className="hero-wrap" style={{ backgroundImage: 'url(images/home_main.webp)', backgroundPosition: '50% 0px' }} data-stellar-background-ratio="0.5">
                <div className="overlay"></div>
                <div className="overlay-2"></div>
                <div className="container">
                    <div className="row no-gutters slider-text justify-content-center align-items-center">
                        <div className="col-lg-8 col-md-6 d-flex align-items-end">
                            <div className="text text-center w-100">
                                <h1 className="mb-4"><strong>Find Rentals</strong> <br />That Save You Money</h1>
                            </div>
                            <div className="mouse">
                                <a href="#properties" className="mouse-icon">
                                    <div className="mouse-wheel"><i className="bi bi-caret-down-fill"></i></div>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
export default Hero;