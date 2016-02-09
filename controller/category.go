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
func (cat *Category) GetCategoryTree(r *http.Request, parentId int64) (CategoryTree, error) {
    //get list of all categories
    categoryModel := model.Category{}
    categories, err := categoryModel.GetAllCategories(r)

    branches := make([]CategoryBranch, 0, 20)

    for _, cat := range categories {
        y:= CategoryBranch {
            Name: cat.Name,
            Id: cat.Id,
            ParentId: cat.ParentId,
        }

        branches = append(branches, y)
    }

    //generate tree
    tree, err := cat.GetCategoryBranch(r, branches, parentId)

    if err != nil {
        return CategoryTree{}, err
    }

    return tree, nil
}

func (cat *Category) GetCategoryBranch(r *http.Request, categories []CategoryBranch, parentId int64) (CategoryTree, error) {

    var finalTree CategoryTree
    var isParent bool = false

    results := make([]CategoryBranch, 0, 20)

    //get initial tree base
    for _, cb := range categories {

        if cb.ParentId == parentId {
            //populate struct
            y:= CategoryBranch {
                Name: cb.Name,
                Id: cb.Id,
                ParentId: cb.ParentId,
            }

            isParent = true

            //get another level down
            cat.GetCategoryBranch(r,categories,cb.Id)

            results = append(results, y)

            log.Println("broheim?!: ", results)
        }
    }

    if isParent == false {
        //
        finalTree.Branch = append(finalTree.Branch, results)

    }

    return finalTree, nil
}
