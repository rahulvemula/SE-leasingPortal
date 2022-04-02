import React, {Component} from "react";
import ReactDom  from "react-dom";
import ReactS3 from 'react-s3';
window.Buffer = window.Buffer || require("buffer").Buffer;


const config = {
    bucketName: 'lease-portal',
    dirName: 'photos', /* optional */
    region: env.a,
    accessKeyId: env.b,
    secretAccessKey:env.c
}


class Upload extends Component{
   
 
    /* this upload function takes creds from aws server and uploads
    the images to the Amazon S3 bucket*/
    upload = (e) => {
        console.log(e.target.files[0]);
        ReactS3.uploadFile(e.target.files[0], config)
            .then((data)=>{
                console.log(data);
                alert("Image uploaded successfully at "+data.location)
            })
            .catch((err)=>{
                console.log(err)
                alert(err);
            })
    }


    render(){
        return(
            <div style={{marginTop : "100px", textAlign: "center"}}>
                <h3>
                    Image Upload tool via Amazon S3
                </h3>
                <input style={{paddingTop: "20px", marginBottom: "20px"}}
                type="file"
                onChange={this.upload}
                />
            </div>
        );
    }
}

export default Upload;

