import Header from "./Header";

function Filters() {
    return (
        <>
        <Header></Header><br/>
            <section className="ftco-section ftco-no-pb">
                <div className="container">
                    <div className="row">
                        <div className="col-md-12">
                            <div className="search-wrap-1">
                                <form action="#" className="search-property-1">
                                    <div className="row">
                                        <div className="col-lg align-items-end">
                                            <div className="form-group">
                                                <label htmlFor="#">Location</label>
                                                <div className="form-field">
                                                    <div className="icon"><i className="bi bi-caret-down-fill"></i></div>
                                                    <input type="text" className="form-control" placeholder="City/Locality Name" />
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-lg align-items-end">
                                            <div className="form-group">
                                                <label htmlFor="#">Property Type</label>
                                                <div className="form-field">
                                                    <div className="select-wrap">
                                                        <div className="icon"><i className="bi bi-caret-down-fill"></i></div>
                                                        <select name="" id="" className="form-control">
                                                            <option value="">Type</option>
                                                            <option value="">Commercial</option>
                                                            <option value="">- Office</option>
                                                            <option value="">Residential</option>
                                                            <option value="">Villa</option>
                                                            <option value="">Condominium</option>
                                                            <option value="">Apartment</option>
                                                        </select>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-lg align-items-end">
                                            <div className="form-group">
                                                <label htmlFor="#">Property Status</label>
                                                <div className="form-field">
                                                    <div className="select-wrap">
                                                        <div className="icon"><i className="bi bi-caret-down-fill"></i></div>
                                                        <select name="" id="" className="form-control">
                                                            <option value="">Type</option>
                                                            <option value="">Rent</option>
                                                            <option value="">Sale</option>
                                                        </select>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-lg align-items-end">
                                            <div className="form-group">
                                                <label htmlFor="#">Price Limit</label>
                                                <div className="form-field">
                                                    <div className="select-wrap">
                                                        <div className="icon"><i className="bi bi-caret-down-fill"></i></div>
                                                        <select name="" id="" className="form-control">
                                                            <option value="">$5,000</option>
                                                            <option value="">$10,000</option>
                                                            <option value="">$50,000</option>
                                                            <option value="">$100,000</option>
                                                            <option value="">$200,000</option>
                                                            <option value="">$300,000</option>
                                                            <option value="">$400,000</option>
                                                            <option value="">$500,000</option>
                                                            <option value="">$600,000</option>
                                                            <option value="">$700,000</option>
                                                            <option value="">$800,000</option>
                                                            <option value="">$900,000</option>
                                                            <option value="">$1,000,000</option>
                                                            <option value="">$2,000,000</option>
                                                        </select>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div className="col-lg align-self-end">
                                            <div className="form-group">
                                                <div className="form-field">
                                                    <input type="submit" value="Search Property" className="form-control btn btn-primary" />
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </>
    )
}

export default Filters;