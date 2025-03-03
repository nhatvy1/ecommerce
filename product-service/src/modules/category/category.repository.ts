import { Injectable } from '@nestjs/common'
import { InjectModel } from '@nestjs/mongoose'
import { Category, CategoryDocument } from './category.schema'
import { Model } from 'mongoose'
import { CreateCategoryDto } from './dto/create_category_dto'

@Injectable()
export class CategoryRepository {
  constructor(
    @InjectModel(Category.name) private readonly categoryModel: Model<Category>
  ) {}

  create(createCategory: CreateCategoryDto) {
    return this.categoryModel.create(createCategory)
  }

  findByIdAndSelect(id: string, select: (keyof Category)[]) {
    return this.categoryModel.find({ _id: id }).select(select).lean()
  }

  update() {}

  delete() {}
}
