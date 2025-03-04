import { Module } from '@nestjs/common'
import { MongooseModule } from '@nestjs/mongoose'
import { Category, CategorySchema } from './category_schema'
import { CategoryController } from './category_controller'
import { CategoryRepository } from './category_repository'
import { CategoryService } from './category_service'

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Category.name, schema: CategorySchema }])
  ],
  controllers: [CategoryController],
  providers: [CategoryRepository, CategoryService]
})
export class CategoryModule {}
