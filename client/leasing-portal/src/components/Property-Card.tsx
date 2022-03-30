function PropertyCard (props: any) {
    return (
        <>
        <div className="property-wrap ">
            <div className="img d-flex align-items-center justify-content-center"
                style={{
                    backgroundImage: `url(images/${props.img}.webp)`,
                }}>
                <div className="list-agent d-flex align-items-center">
                    <a href="/lease-confirmation" className="agent-info d-flex align-items-center">
                        <div
                            className="img-2 rounded-circle"
                            style={{backgroundImage:"url(images/xperson_1.jpg.pagespeed.ic.a2MnMHMs44.webp)",}}
                        ></div>
                    </a>
                </div>
            </div>
            <div className="text">
                <p className="price mb-3">
                    <span className="orig-price">
                        ${props.rent}<small>/mo</small>
                    </span>
                </p>
                <h3 className="mb-0">
                    <a href="properties-single.html">{props.label}</a>
                </h3>
                <ul className="property_list">
                    <li>
                        <span><i className="bi bi-house-fill"></i></span>
                        {props.sqft} sqft
                    </li>
                </ul>
            </div>
        </div>
        </>
    )
}

export default PropertyCard;