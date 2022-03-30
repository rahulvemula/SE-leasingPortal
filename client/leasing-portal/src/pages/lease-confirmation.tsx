import "../index.css";
import html2canvas from "html2canvas";
import {jsPDF} from "jspdf";
// import InvoiceGenerator from "../utils/leaseDocGenerator";

function Done() {

    let exportPdf = () => {
        let t : HTMLElement = document.querySelector(".msg") as HTMLElement;
        html2canvas(t).then(canvas => {
           document.body.appendChild(canvas);  // if you want see your screenshot in body.
           const imgData = canvas.toDataURL('image/png');
           const pdf = new jsPDF("l");
           pdf.addImage(imgData, 'PNG', 0, 0, 300, 200);
           pdf.save("download.pdf"); 
       });
   
    }


  return (
      <div>
    <div className="msg">
      <h3>Congratulations! Welcome to the society!</h3>

      <div>&nbsp;&nbsp;Please find the lease details below</div>

      <div className="grid-container ">
        <div className="grid-item"><strong>Name</strong></div>
        <div className="grid-item">Lahari</div>

        <div className="grid-item">Email</div>
        <div className="grid-item">lbarad@blahblah.com</div>

        <div className="grid-item">Property Name</div>
        <div className="grid-item">The Niche Student Appts</div>

        <div className="grid-item">Lease start date</div>
        <div className="grid-item">03/29/2022</div>

        <div className="grid-item">Lease end date</div>
        <div className="grid-item">03/29/2023</div>
      </div>
      <br />
      
    </div>
    <div className="text-center">
        <button className="btn btn-primary" onClick={exportPdf}>Print</button>
      </div>
      <br />
    </div>
  );
}

export default Done;
