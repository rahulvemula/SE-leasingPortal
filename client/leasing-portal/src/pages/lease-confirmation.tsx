import "../index.css";
import { jsPDF } from "jspdf";
import { useSelector } from "react-redux";

interface Details {
  [key: string]: any;
}

function Done() {
  let message = "Congratulations! Welcome to the society!";
  let detailText = "Please find the lease details below";
  const leaseData = useSelector((state: any) => state.leaseData);
  
  let details: Details = {
    Name: leaseData.name,
    Email: leaseData.email,
    "Property Name": "The Niche Student Appts",
    "Lease start date": leaseData.startDate,
    "Lease end date": leaseData.endDate,
  };

  let exportPdf = () => {
    var doc = new jsPDF("p", "pt", "letter");
    doc.text("HOUS.", 40, 40).setFontSize(20).setFont("bold");

    doc.text(message, 150, 100).setFontSize(15);
    doc.setFontSize(12);
    doc.text(detailText, 150, 150);

    let topPadding = 200;
    Object.keys(details).forEach((key) => {
      doc.text(key, 150, topPadding).setFont("bold");
      doc.text(details[key], 300, topPadding);
      topPadding = topPadding + 50;
    });

    doc.save("Lease-Document.pdf");
  };

  return (
    <div>
      <div className="msg">
        <h3 id="confirmation-msg">{message}</h3>

        <div>&nbsp;&nbsp;{detailText}</div>
        <div className="grid-container ">
          {Object.keys(details).map((key) => {
            return (
              <>
                <div className="grid-item">
                  <strong>{key}</strong>
                </div>
                <div className="grid-item">{details[key]}</div>
              </>
            );
          })}
        </div>
        <br />
      </div>
      <div className="text-center">
        <button id="print-doc" className="btn btn-primary" onClick={exportPdf}>
          Print Document
        </button>
      </div>
      <br />
    </div>
  );
}

export default Done;
