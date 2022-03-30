// const PDFGenerator = require('pdfkit')
// const fs = require('fs')

// export default function InvoiceGenerator (invoice){

//     function generateHeaders(doc) {
//         const billingAddress = invoice.addresses.billing

//         doc
//             .image('./door-company-logo.jpg', 0, 0, { width: 250})
//             .fillColor('#000')
//             .fontSize(20)
//             .text('INVOICE', 275, 50, {align: 'right'})
//             .fontSize(10)
//             .text(`Invoice Number: ${invoice.invoiceNumber}`, {align: 'right'})
//             .text(`Due: ${invoice.dueDate}`, {align: 'right'})
//             .text(`Balance Due: $${invoice.subtotal - invoice.paid}`, {align: 'right'})
//             .moveDown()
//             .text(`Billing Address:\n ${billingAddress.name}\n${billingAddress.address}\n${billingAddress.city}\n${billingAddress.state},${billingAddress.country}, ${billingAddress.postalCode}`, {align: 'right'})
    
//         const beginningOfPage = 50
//         const endOfPage = 550

//         doc.moveTo(beginningOfPage,200)
//             .lineTo(endOfPage,200)
//             .stroke()
                
//         doc.text(`Memo: ${invoice.memo || 'N/A'}`, 50, 210)

//         doc.moveTo(beginningOfPage,250)
//             .lineTo(endOfPage,250)
//             .stroke()

//     }

//     function generateTable(doc) {
//         const tableTop = 270
//         const itemCodeX = 50
//         const descriptionX = 100
//         const quantityX = 250
//         const priceX = 300
//         const amountX = 350

//         doc
//             .fontSize(10)
//             .text('Item Code', itemCodeX, tableTop, {bold: true})
//             .text('Description', descriptionX, tableTop)
//             .text('Quantity', quantityX, tableTop)
//             .text('Price', priceX, tableTop)
//             .text('Amount', amountX, tableTop)

//         const items = invoice.items
//         let i = 0


//         for (i = 0; i < items.length; i++) {
//             const item = items[i]
//             const y = tableTop + 25 + (i * 25)

//             doc
//                 .fontSize(10)
//                 .text(item.itemCode, itemCodeX, y)
//                 .text(item.description, descriptionX, y)
//                 .text(item.quantity, quantityX, y)
//                 .text(`$ ${item.price}`, priceX, y)
//                 .text(`$ ${item.amount}`, amountX, y)
//         }
//     }

//     function generateFooter(doc) {
//         doc
//             .fontSize(10)
//             .text(`Payment due upon receipt. `, 50, 700, {
//                 align: 'center'
//             })
//     }

//     function generate(invoice) {

//         console.log(invoice)

//         const fileName = `Invoice ${invoice.invoiceNumber}.pdf`

//         // pipe to a writable stream which would save the result into the same directory
//         theOutput.pipe(fs.createWriteStream(fileName))

//         generateHeaders(theOutput)

//         theOutput.moveDown()

//         generateTable(theOutput)

//         generateFooter(theOutput)
        

//         // write out file
//         theOutput.end()

//     }

//     generate(invoice);
// }

// InvoiceGenerator