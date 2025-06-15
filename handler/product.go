package handler

import (
	"aia_backend/models"
	"aia_backend/utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/alioygur/gores"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func UploadProductFileHandler(rw http.ResponseWriter, rq *http.Request) {
	ctx := rq.Context()
	param, file, err := extraPutCodeMultipartParams(rq)
	if err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, err)
		return
	}
	if param["category"] == "" {
		_ = gores.JSON(rw, http.StatusBadRequest, fmt.Errorf("should spcify category"))
		return
	}
	if param["filename"] == "" {
		_ = gores.JSON(rw, http.StatusBadRequest, fmt.Errorf("should spcify filename"))
		return
	}

	objectKey := fmt.Sprintf("products/%s", utils.RandomStringWithLength(10))
	// 储存 file 到 tos 上
	if err = upload(ctx, objectKey, file); err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Errorf("should spcify filename"))
		return
	}

	p := models.ProductFile{
		Category:   models.ProductType(param["category"]),
		Name:       param["filename"],
		ObjectKeys: objectKey,
	}

	if err = models.SaveProducts(ctx, p); err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, err)
		return
	}

	_ = gores.JSON(rw, http.StatusOK, p)
	return
}

func extraPutCodeMultipartParams(
	rq *http.Request,
) (params map[string]string, file io.ReadCloser, err error) {
	params = map[string]string{
		"zip_file_size": "",
	}

	// Parse multipart form data.
	reader, err := rq.MultipartReader()
	if err != nil {
		return
	}

	// The incoming parameters.
	b := bytes.Buffer{}

	// Parse the request form.
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if part == nil {
			err = fmt.Errorf("failed to read form part, %v", err)
			return nil, nil, err
		}

		if part.FormName() == "file" {
			file = part
			break
		}

		b.Reset()
		_, err = b.ReadFrom(part)
		if err != nil {
			return nil, nil, err
		}

		if _, ok := params[part.FormName()]; ok {
			params[part.FormName()] = b.String()
		}
	}
	if file == nil {
		err = errors.New("parameter zip_file should be specified")
		return
	}
	return
}

func DownloadProductFileHandler(rw http.ResponseWriter, rq *http.Request) {
	ctx := rq.Context()
	productID := mux.Vars(rq)["product_id"]
	if productID == "" {
		_ = gores.JSON(rw, http.StatusBadRequest, fmt.Errorf("should spcify product_id"))
		return
	}

	product, err := models.GetProductByID(ctx, productID)
	if err != nil {
		_ = gores.JSON(rw, http.StatusBadRequest, fmt.Errorf("find product failed, err: %v", err))
		return
	}

	objectKey := product.ObjectKeys
	res, err := download(ctx, objectKey)
	if err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Errorf("download failed, err: %v", err))
		return
	}

	_, err = io.Copy(rw, res)
	if err != nil {
		_ = gores.JSON(rw, 500, fmt.Errorf("copy failed, err: %v", err))
		return
	}
}

func ListProductsHandler(rw http.ResponseWriter, rq *http.Request) {
	ctx := rq.Context()

	p, err := models.ListProducts(ctx)
	if err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, err)
		return
	}

	_ = gores.JSON(rw, http.StatusOK, p)
}
