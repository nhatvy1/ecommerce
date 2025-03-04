import { Injectable, NotFoundException } from '@nestjs/common'
import { CategoryRepository } from './category.repository'
import { ErrorMessage } from 'src/utils/response'
import { CreateCategoryDto, SaveCategoryDto } from './dto/create_category_dto'
import { plainToInstance } from 'class-transformer'
import { UpdateCategoryDto } from './dto/update_category_dto'

@Injectable()
export class CategoryService {
  constructor(private readonly categoryRepository: CategoryRepository) {}

  async findById(id: string) {
    const category = await this.categoryRepository.findByIdAndSelect(id, [
      'name'
    ])
    if (!category) throw new NotFoundException(ErrorMessage.ErrNotFound)
  }

  async create(body: CreateCategoryDto) {
    const { name, parentId } = body

    let rightValue: any

    if (parentId) {
      const parentCategory = await this.categoryRepository.findById(parentId)
      if (!parentCategory) {
        throw new NotFoundException(ErrorMessage.ErrNotFound)
      }

      rightValue = parentCategory.right

      await this.categoryRepository.updateMany({
        filter: { right: { $gte: rightValue } },
        update: { right: 2 }
      })

      await this.categoryRepository.updateMany({
        filter: { right: { $gt: rightValue } },
        update: { right: 2 }
      })
    } else {
      const maxRightValue = await this.categoryRepository.getMaxRightValue()
      if (maxRightValue) {
        rightValue = maxRightValue.right + 1
      } else {
        rightValue = 1
      }
    }

    const categoryRight = rightValue + 1
    const categoryLeft = rightValue
    const dataCreate = plainToInstance(SaveCategoryDto, {
      name,
      left: categoryLeft,
      right: categoryRight,
      parentId: parentId
    })

    await this.categoryRepository.create(dataCreate)

    return 1
  }

  async update(id: string, body: UpdateCategoryDto) {
    const category = await this.categoryRepository.update({
      filter: { _id: id },
      update: { body }
    })
    if (!category) {
      throw new NotFoundException(ErrorMessage.ErrNotFound)
    }
    return 1
  }

  async delete(id: string) {
    const category = await this.categoryRepository.findById(id)
    if (!category) {
      throw new NotFoundException(ErrorMessage.ErrNotFound)
    }

    const leftValue = category.left
    const rightValue = category.right

    const width = rightValue - leftValue + 1

    // delete all child category
    await this.categoryRepository.deleteMany({
      left: { $gte: leftValue, $lte: rightValue }
    })

    // update left and right values
    await this.categoryRepository.updateMany({
      filter: { right: { $gt: rightValue } },
      update: { $inc: { right: -width } }
    })

    await this.categoryRepository.updateMany({
      filter: { left: { $gt: rightValue } },
      update: { $inc: { left: -width } }
    })

    return 1
  }
}
