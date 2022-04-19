import  './AboutUs.css'
function AboutUs() {
    return (<>
  
 
 <h2 style={{textAlign:'center', marginTop:'100px'}}>Our Team</h2>
 <div  className="aboutUsRow">
   <div  className="column">
     <div className="card">
       <img     src="https://avatars.githubusercontent.com/u/98243209?v=4" alt="Lahari" style={{width:'100%'}}/>
       <div className="container">
         <h2>Lahari Barad</h2>
         <p className="title">React developer</p>
         <p>laharibarad@ufl.edu</p>
         <p><button className="button">Contact</button></p>
       </div>
     </div>
   </div>
 
   <div className="column">
     <div className="card">
       <img src="https://avatars.githubusercontent.com/u/8132049?v=4" alt="Rahul" style={{width:'100%'}}/>
       <div className="container">
         <h2>Rahul Vemula</h2>
         <p className="title">React developer</p>
         <p>rahulvemula@ufl.edu</p>
         <p><button className="button">Contact</button></p>
       </div>
     </div>
   </div>
 
   <div className="column">
     <div className="card">
       <img src="https://avatars.githubusercontent.com/u/34827172?v=4" alt="Mitul" style={{width:'100%'}}/>
       <div className="container">
         <h2>Mitul Mandaliya</h2>
         <p className="title">Go lang developer</p>
         <p>mitulmandaliya@ufl.edu</p>
         <p><button className="button">Contact</button></p>
       </div>
     </div>
   </div>
   <div className="column">
     <div className="card">
       <img src="https://avatars.githubusercontent.com/u/16957347?s=400&u=4da3a5e6d98a8bbc42f3ee4389c43cae1808f2fe&v=4" alt="Vamsi" style={{width:'100%'}}/>
       <div className="container">
         <h2>Vamsi Viswanath</h2>
         <p className="title">Go lang developer</p>
         <p>vbethmasetty@ufl.edu</p>
         <p><button className="button">Contact</button></p>
       </div>
     </div>
   </div>
   
 </div>
 
  
    </>)
 }
 
 export default AboutUs