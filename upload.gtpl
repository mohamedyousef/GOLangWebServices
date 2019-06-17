<html>
  <head>
         <title>Test Upload a File</title>
  </head>
  <body>
<form enctype="multipart/form-data" action="http://127.0.0.1:5000/upload" method="post">
          {{/* 1. File input */}}
          <input type="text" name="pain" />
          <input type="text" name="h" />
          <input type="text" name="blood" />
          <input type="text" name="age" />
          <input type="file" name="uploadfile" />

          {{/* 2. Submit button */}}

          <input type="submit" value="upload file" />

      </form>

  </body>
  </html>