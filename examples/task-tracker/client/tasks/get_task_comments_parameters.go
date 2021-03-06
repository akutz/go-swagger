package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetTaskCommentsParams creates a new GetTaskCommentsParams object
// with the default values initialized.
func NewGetTaskCommentsParams() *GetTaskCommentsParams {
	var (
		pageSizeDefault int32 = int32(20)
	)
	return &GetTaskCommentsParams{
		PageSize: &pageSizeDefault,
	}
}

/*GetTaskCommentsParams contains all the parameters to send to the API endpoint
for the get task comments operation typically these are written to a http.Request
*/
type GetTaskCommentsParams struct {

	/*ID
	  The id of the item

	*/
	ID int64
	/*PageSize
	  Amount of items to return in a single page

	*/
	PageSize *int32
	/*Since
	  The created time of the oldest seen comment

	*/
	Since *strfmt.DateTime
}

// WithID adds the id to the get task comments params
func (o *GetTaskCommentsParams) WithID(id int64) *GetTaskCommentsParams {
	o.ID = id
	return o
}

// WithPageSize adds the pageSize to the get task comments params
func (o *GetTaskCommentsParams) WithPageSize(pageSize *int32) *GetTaskCommentsParams {
	o.PageSize = pageSize
	return o
}

// WithSince adds the since to the get task comments params
func (o *GetTaskCommentsParams) WithSince(since *strfmt.DateTime) *GetTaskCommentsParams {
	o.Since = since
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskCommentsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince strfmt.DateTime
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince.String()
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
