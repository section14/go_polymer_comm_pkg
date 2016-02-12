package controller

import (
    "net/http"
    "encoding/json"
    "strconv"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Category struct {
    Name string `json:"name"`
    ParentId int64 `json:"parentid"`
}

type CategoryReturn struct {
    Name string
    Id int64
    ParentId int64
    Products []int64
    Key string
}

type CategoryBranch struct {
    Name string
    Id int64
    ParentId int64
    SubBranch []CategoryBranch
}

type CategoryTree struct {
    Branch [][]CategoryBranch
}

func (cat *Category) CreateCategory(r *http.Request) (bool, error) {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&cat)

    if err != nil {
        //handle err
        return false, err
    }

    //populate category data
    categoryModel := model.Category{}
    categoryModel.Name = cat.Name
    err = categoryModel.CreateCategory(r, cat.ParentId)

    if err != nil {
        log.Println(err)
    }

    return true, nil
}

func (cat *Category) GetCategories(r *http.Request) ([]CategoryReturn, error) {
    parentCat := r.Header.Get("Parent-Category")

    //convert header string to int64
    parentId, err := strconv.ParseInt(parentCat, 10, 64)

    if err != nil {
        return []CategoryReturn{}, err
    }

    categoryModel := model.Category{}
    categories, err := categoryModel.GetCategories(r, parentId)

    if err != nil {
        log.Println(err)
    }

    results := make([]CategoryReturn, 0, 20)

    for _, r := range categories {
        y := CategoryReturn {
            Name: r.Name,
            Id: r.Id,
            ParentId: r.ParentId,
            Products: r.Products,
            Key: r.Key,
        }

        results = append(results, y)
    }

    return results, nil
}

func (cat *Category) GetAllCategories(r *http.Request) ([]CategoryReturn, error) {
    categoryModel := model.Category{}
    categories, err := categoryModel.GetAllCategories(r)

    if err != nil {
        log.Println(err)
    }

    results := make([]CategoryReturn, 0, 20)

    for _, r := range categories {
        y := CategoryReturn {
            Name: r.Name,
            Id: r.Id,
            ParentId: r.ParentId,
            Products: r.Products,
            Key: r.Key,
        }

        results = append(results, y)
    }

    return results, nil
}

//this returns all the categories in a structered json tree
func (cat *Category) GetCategoryTree(r *http.Request, parentId int64) ([]CategoryBranch, error) {
    //get list of all categories
    categoryModel := model.Category{}
    categories, err := categoryModel.GetAllCategories(r)

    if err != nil {
        return []CategoryBranch{}, err
    }

    branches := make([]CategoryBranch, 0, 20)

    for _, cat := range categories {
        y:= CategoryBranch {
            Name: cat.Name,
            Id: cat.Id,
            ParentId: cat.ParentId,
        }

        branches = append(branches, y)
    }

    categoryTree, err := cat.GetCategoryBranch(branches, parentId)

    if err != nil {
        return []CategoryBranch{}, err
    }

    return categoryTree, nil
}

func (cat *Category) GetCategoryBranch(categories []CategoryBranch, parentId int64) ([]CategoryBranch, error) {

    branch := make([]CategoryBranch, 0, 20)

    //get initial tree base
    for _, cs := range categories {

        if cs.ParentId == parentId {
            newBranch, err := cat.GetCategoryBranch(categories, cs.Id)

            if err != nil {
                log.Println("something in GetCategoryBranch")
            }

            var y CategoryBranch

            if newBranch != nil {
                log.Println("has children: ", newBranch)

                y = CategoryBranch {
                    Name: cs.Name,
                    Id: cs.Id,
                    ParentId: cs.ParentId,
                    SubBranch: newBranch,
                }
            } else {
                y = CategoryBranch {
                    Name: cs.Name,
                    Id: cs.Id,
                    ParentId: cs.ParentId,
                }
            }

            branch = append(branch, y)
        }

    }

    return branch, nil
}

func (cat *Category) GetUniqueParents(categories []CategoryBranch) []int64{
    parentList := make([]int64, 0, 20)

    for _, cb := range categories {
        newParent := AddUniqueParent(cb.ParentId, parentList)

        if newParent == true {
            parentList = append(parentList, cb.ParentId)
        }
    }

    return parentList
}

func AddUniqueParent(id int64, parentList []int64) bool {
    for _, cb := range parentList {
        if id == cb {
            return false
        }
    }

    return true
}
