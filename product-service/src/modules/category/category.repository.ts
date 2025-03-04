import { Injectable } from '@nestjs/common'
import { InjectModel } from '@nestjs/mongoose'
import { Category } from './category.schema'
import { FilterQuery, Model, QueryOptions, UpdateQuery } from 'mongoose'
import { SaveCategoryDto } from './dto/create_category_dto'

@Injectable()
export class CategoryRepository {
  constructor(
    @InjectModel(Category.name) private readonly categoryModel: Model<Category>
  ) {}

  create(createCategory: SaveCategoryDto) {
    return this.categoryModel.create(createCategory)
  }

  findById(id: string) {
    return this.categoryModel.findById(id).lean()
  }

  findByIdAndSelect(id: string, select: (keyof Category)[]) {
    return this.categoryModel.find({ _id: id }).select(select).lean()
  }

  async getMaxRightValue() {
    return this.categoryModel
      .findOne({}, 'right', {
        sort: { right: -1 }
      })
      .lean()
  }

  updateMany({
    filter,
    update
  }: {
    filter: FilterQuery<Category>
    update: UpdateQuery<Category>
  }) {
    return this.categoryModel.updateMany(filter, update)
  }

  update({
    filter,
    update,
    options
  }: {
    filter: FilterQuery<Category>
    update: UpdateQuery<Category>,
    options?: QueryOptions<Category>
  }) {
    return this.categoryModel.findOneAndUpdate(filter, update, options).lean()
  }

  deleteMany(filter: FilterQuery<Category>) {
    return this.categoryModel.deleteMany(filter)
  }
}
