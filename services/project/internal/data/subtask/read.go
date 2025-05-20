package subtaskdata

import (
	"context"

	"pms.pkg/type/list"
)

func (r *Repository) List(ctx context.Context, filter SubTaskFilter) (res list.List[SubTask], err error) {

	return res, err
}
