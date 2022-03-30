import PropertyCard from "../components/Property-Card"

function ApptListing() {
    let appts = [
        {
            label: 'Full house - 4bhk',
            rent: 2509,
            sqft: 2400,
            img: 'House1',
            id: 1
    
        },
        {
            label: 'Full house - 3bhk',
            rent: 1500,
            sqft: 1800,
            img: 'House2',
            id: 2,
    
        },
        {
            label: 'Single room in 3 Bed 2 Bath (Shared)',
            rent: 509,
            sqft: 400,
            img: '3bhk',
            id: 3
    
        },
        {
            label: '3 Bed 2 Bath (Private)',
            rent: 549,
            sqft: 400,
            img: '3bhk',
            id:4
    
        },
        {
            label: '4 Bed 3 Bath (Shared)',
            rent: 449,
            sqft: 350,
            img: '4bhk',
            id:5
    
        },
        {
            label: '4 Bed 3 Bath (Private)',
            rent: 489,
            sqft: 350,
            img: '4bhk',
            id:6
    
        },
        {
            label: '4 Bed 4 Bath (Personal)',
            rent: 469,
            sqft: 350,
            img: '4bhkP',
            id:7
    
        },
        {
            label: '4 Bed 4 Bath (Private)',
            rent: 469,
            sqft: 350,
            img: '4bhkP',
            id:8
    
        }



    ];

    let rowSize = 3;

    function getCols() {
        let cols = [];
        while(appts.length > 0) {
            cols.push(appts.splice(0, rowSize));
        }
        return cols;
    }
    

    return (
        <>
            <section className="ftco-section goto-here">
                <div className="container">
                    <br/>
                    {
                        getCols().map(function(col) {
                            return (
                                <div className="row">
                                    {
                                        col.map(function (i) {
                                            return (
                                                <div className="col-md-4">
                                                    <PropertyCard {...i}></PropertyCard>
                                                </div>
                                            )
                                        })
                                    }
                                </div> 
                            )
                        })
                    }

                </div>
            </section>
        </>
    )
}

export default ApptListing;